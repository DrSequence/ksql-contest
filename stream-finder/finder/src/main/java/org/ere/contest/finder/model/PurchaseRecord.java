package org.ere.contest.finder.model;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public record PurchaseRecord(
        @JsonProperty("order_id") String orderId,
        @JsonProperty("user_id") String userId,
        long orderTime,
        @JsonProperty("total_price") double totalPrice,
        List<ProductRecord> products
) {}