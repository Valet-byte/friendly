package com.valet.gurren.repo.cache.messageTokenCache

import com.valet.gurren.model.MessagingToken

interface MessagingTokenCache {

    fun remove(username: String)
    fun save(messagingToken: MessagingToken)
    fun exist(username: String): Boolean
    fun getMessagingToken(username: String): MessagingToken
}