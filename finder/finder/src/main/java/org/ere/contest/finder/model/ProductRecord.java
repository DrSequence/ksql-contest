package org.ere.contest.finder.model;

import com.fasterxml.jackson.annotation.JsonProperty;

public record ProductRecord(
        @JsonProperty("product_id") String productId,
        String category,
        double price
) {}