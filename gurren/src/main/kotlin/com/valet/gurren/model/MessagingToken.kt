package com.valet.gurren.model

import jakarta.persistence.Entity
import jakarta.persistence.Id

@Entity
data class MessagingToken (
    @Id val username: String = "",
    val token: String = ""
    )