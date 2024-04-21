package org.ere.contest.finder.service;

import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.streams.KeyValue;
import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.kstream.Consumed;
import org.apache.kafka.streams.kstream.KStream;
import org.apache.kafka.streams.kstream.Printed;
import org.apache.kafka.streams.kstream.Produced;
import org.ere.contest.finder.config.prop.StreamProperty;
import org.ere.contest.finder.model.EnrichedPurchase;
import org.ere.contest.finder.model.EventRecord;
import org.ere.contest.finder.model.ProductRecord;
import org.ere.contest.finder.model.PurchaseRecord;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.annotation.EnableKafka;
import org.springframework.kafka.annotation.EnableKafkaStreams;
import org.springframework.kafka.support.serializer.JsonSerde;

import java.util.Optional;

@Configuration
@EnableKafka
@EnableKafkaStreams
public class KafkaStreamsConfig {

    private final StreamProperty streamProperty;

    public KafkaStreamsConfig(StreamProperty streamProperty) {
        this.streamProperty = streamProperty;
    }

    @Bean
    public KStream<String, EnrichedPurchase> kStream(StreamsBuilder kStreamBuilder) {
        KStream<String, PurchaseRecord> shopStream = kStreamBuilder.stream(
                streamProperty.shopTopic(),
                Consumed.with(Serdes.String(), new JsonSerde<>(PurchaseRecord.class))
        );

        KStream<String, EventRecord> viewStream = kStreamBuilder.stream(
                streamProperty.eventTopic(),
                Consumed.with(Serdes.String(), new JsonSerde<>(EventRecord.class))
        );

        KStream<String, EnrichedPurchase> enrichedPurchases = shopStream.join(
                viewStream.toTable(),
                (purchase, view) -> {
                    if (purchase.userId().equals(view.userId())) {
                        Optional<String> category = purchase.products()
                                .stream()
                                .map(ProductRecord::category)
                                        .filter(f -> f.equals(view.category()))
                                .findFirst();

                        if (category.isPresent()) {
                            return new EnrichedPurchase(view, purchase);
                        }
                    }
                    return null;
                }
        ).map(KeyValue::new)
                .filter((i, v) -> v != null);

        enrichedPurchases.print(Printed.toSysOut());
        enrichedPurchases.to(
                streamProperty.resultTopic(),
                Produced.with(Serdes.String(), new JsonSerde<>(EnrichedPurchase.class))
        );

        return enrichedPurchases;
    }
}