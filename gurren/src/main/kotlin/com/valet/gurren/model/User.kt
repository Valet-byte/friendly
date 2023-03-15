package com.valet.gurren.model

import jakarta.persistence.Entity
import jakarta.persistence.Id
import java.time.Instant
import java.util.Date

@Entity
data class User(
    @Id val username: String = "",
    val dateOfBirth: Date = Date.from(Instant.now()),
    val isPremium: Boolean = false,
    val isEnable: Boolean = true,
    val city: String = "Москва",
    val messagingToken: String = ""
){
    companion object{
        fun createUser(userData: UserData): User {
            return User(username = userData.username,
                        dateOfBirth = userData.date,
                        city = userData.city,
                        messagingToken = userData.messagingToken)
        }
    }
}

data class UserData(
    val username: String,
    val date: Date,
    val city: String,
    val messagingToken: String
)