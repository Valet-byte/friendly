package com.valet.gurren.model

import java.util.*

data class Message(
    var id: String = "",
    val fromUser: String,
    val body: String, /* if type == PICTURE */
    val type: MessageType,
    val date: Date,
    val chatId: String
)

enum class MessageType{
    TEXT, PICTURE, DOCUMENT
}