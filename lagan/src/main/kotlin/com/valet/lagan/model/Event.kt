package com.valet.lagan.model

import java.util.*

data class Event(
    private val id: UUID,
    private val usernameCreator: String,
    private val name: String,
    private val address: String,
    private val photoLink: String,
    private val peopleMaxSize: UInt,
    private val peopleMinSize: UInt,
    private val peopleSize: UInt,
    private val peopleFriendSize: UInt,
    private val startTime: Date,
    private val endTime: Date,
    private val accessType: EventType
)

enum class EventType {
    ALL_USER, FRIEND_ONLY, BY_INVITATION, BY_INVITATION_OR_REQUEST
}
