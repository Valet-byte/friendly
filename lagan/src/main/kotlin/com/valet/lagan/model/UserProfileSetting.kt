package com.valet.lagan.model

data class UserProfileSetting(
    private val username: String,
    private val isCloseProfile: Boolean,
    private val messageFromStrangers: Boolean,
    private val isHideMode: Boolean
)
