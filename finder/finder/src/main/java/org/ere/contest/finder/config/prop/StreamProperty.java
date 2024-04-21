package org.ere.contest.finder.config.prop;

import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "app.streams")
public record StreamProperty(
        String shopTopic,
        String eventTopic,
        String resultTopic
) {}
