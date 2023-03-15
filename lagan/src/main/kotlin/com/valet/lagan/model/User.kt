package com.valet.lagan.model

import java.util.Date

data class User(
    private val username: String,
    private val dateOfBirth: Date,
    private val isPremium: Boolean,
    private val isEnable: Boolean,
    private val city: String,
    private val serialNumber: String
)

data class UserData(
    private val username: String,
    private val date: Date,
    private val city: String,
    private val serialNumber: String,
    private val messagingToken: String
)