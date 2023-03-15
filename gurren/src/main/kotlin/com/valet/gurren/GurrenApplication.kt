package com.valet.gurren

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.Bean
import org.springframework.data.redis.core.RedisTemplate

@SpringBootApplication
class GurrenApplication {

}

fun main(args: Array<String>) {
    runApplication<GurrenApplication>(*args)
}
