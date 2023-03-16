package com.valet.gurren.model

import jakarta.persistence.Entity
import jakarta.persistence.Id

@Entity
data class Description(
    @Id private val username: String = "",
    private val body: String = ""/* max 200 symbols */
)
