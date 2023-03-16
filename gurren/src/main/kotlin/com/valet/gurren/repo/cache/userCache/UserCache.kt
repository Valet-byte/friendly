package com.valet.gurren.repo.cache.userCache

import com.valet.gurren.model.User

interface UserCache {
    fun save(user: User, userId: String) : Int
    fun delete(userId: String) : Long
    fun getUser(userId: String) : User
    fun hasUser(userId: String) : Boolean
}