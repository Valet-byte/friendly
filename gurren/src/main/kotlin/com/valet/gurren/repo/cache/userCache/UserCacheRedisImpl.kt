package com.valet.gurren.repo.cache.userCache

import com.valet.gurren.model.User
import org.springframework.beans.factory.annotation.Value
import org.springframework.data.redis.core.RedisTemplate
import org.springframework.stereotype.Service
import java.time.Duration

@Service
class UserCacheRedisImpl(private val redisTemplate: RedisTemplate<String, String>,
                         @Value("\${cache.duration-time}") private val  durationTime: Long) : UserCache {

    override fun save(user: User, userId: String): Int {
        val hashOperation = redisTemplate.opsForHash<String, String>()
        hashOperation.put("users", userId, user.toString())
        redisTemplate.expire("users", Duration.ofMinutes(durationTime))
        return 1

    }

    override fun delete(userId: String): Long {
        val hashOperation = redisTemplate.opsForHash<String, String>()
        val result = hashOperation.delete("users", userId)
        redisTemplate.expire("users", Duration.ofMinutes(durationTime))
        return result
    }

    override fun getUser(userId: String): User {
        val hashOperation = redisTemplate.opsForHash<String, User>()
        val result = hashOperation.get("users", userId)
        redisTemplate.expire("users", Duration.ofMinutes(durationTime))
        return result!!
    }

    override fun hasUser(userId: String): Boolean = redisTemplate.opsForHash<String, String>().hasKey("users", userId)
}