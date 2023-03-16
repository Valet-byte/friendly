package com.valet.gurren.model

import jakarta.persistence.Entity
import jakarta.persistence.Id
import java.time.Instant
import java.util.Date

@Entity
data class User(
    @Id var username: String = "",
    var dateOfBirth: Date = Date.from(Instant.now()),
    var isPremium: Boolean = false,
    var isEnable: Boolean = true,
    var city: String = "Москва"
){
    companion object{
        fun createUser(userData: UserData): User {
            return User(dateOfBirth = userData.date,
                        city = userData.city)
        }
    }
}

data class UserData(
    val date: Date,
    val city: String,
    val messagingToken: String
)