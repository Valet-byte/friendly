package com.valet.gurren.model

import jakarta.persistence.Entity
import jakarta.persistence.Id

@Entity
data class UserProfileSetting(
    @Id val username: String = "",
    val isCloseProfile: Boolean = false,
    val messageFromStrangers: Boolean = true,
    val isHideMode: Boolean = false
)
