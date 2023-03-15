package com.valetbyte.gurren.model

import java.util.UUID

data class Chat(
    private val id: UUID,
    private val username1: String,
    private val username2: String
)
