package com.valet.gurren.repo.cache.messageTokenCache

import com.valet.gurren.model.MessagingToken
import org.springframework.data.redis.core.RedisTemplate
import org.springframework.stereotype.Service

@Service
class MessagingTokenCacheRedisImpl(private val redisTemplate: RedisTemplate<String, String>) : MessagingTokenCache {
    override fun remove(username: String) {
        redisTemplate.opsForHash<String, String>().delete("message-token", username)
    }

    override fun save(messagingToken: MessagingToken) {
        redisTemplate.opsForHash<String, String>().put("message-token", messagingToken.username, messagingToken.token)
    }

    override fun exist(username: String): Boolean {
        return redisTemplate.opsForHash<String, String>().hasKey("message-token", username)
    }

    override fun getMessagingToken(username: String): MessagingToken {
        val messagingToken: MessagingToken

        val token = redisTemplate.opsForHash<String, String>().get("message-token", username)

        messagingToken = MessagingToken(username, token!!)
        return messagingToken
    }
}