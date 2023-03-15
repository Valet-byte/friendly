package com.valetbyte.gurren.model

data class SocialNetwork(
    private val username: String,
    private val vkLink: String? = null,
    private val telegramLink: String? = null,
    private val instagramLink: String? = null,
    private val tiktokLink: String? = null,
    private val youtubeLink: String? = null
)
