package org.ere.contest.finder;

import org.ere.contest.finder.config.prop.StreamProperty;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;

@SpringBootApplication
@EnableConfigurationProperties(StreamProperty.class)
public class FinderApplication {

    public static void main(String[] args) {
        SpringApplication.run(FinderApplication.class, args);
    }

}
