package com.valetbyte.gurren.model

import java.util.*

data class Message(
    private val fromUser: String,
    private val body: String, /* if type == PICTURE */
    private val type: MessageType,
    private val date: Date,
    private val chatId: UUID
)

enum class MessageType{
    TEXT, PICTURE, DOCUMENT
}