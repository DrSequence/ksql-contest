package org.ere.contest.finder.model;

import com.fasterxml.jackson.annotation.JsonProperty;

public record EventRecord(
        @JsonProperty("product_id") String productId,
        long timestamp,
        String category,
        String app,
        @JsonProperty("session_id") String sessionId,
        @JsonProperty("user_id") String userId
) {}