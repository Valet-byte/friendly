package com.valet.gurren.service.userService.sesionId


import com.google.firebase.database.FirebaseDatabase
import com.google.firebase.messaging.Message
import org.springframework.stereotype.Service

@Service
class SessionIdServiceFirebaseRealtimeDatabaseImpl(
    private val firebaseDatabase: FirebaseDatabase
) : SessionIdService {
    override fun updateSessionId(uid: String, username: String,  sessionId: String) {

        val mess = Message.builder().putData("username", username)
            .putData("sessionId", sessionId)
            .setTopic(uid)
            .build()

        firebaseDatabase.reference.child(uid).setValue(UserAccountInfo(username, sessionId)) { _, _ ->

        }
    }

    data class UpdateData(
        val userAccountInfo: UserAccountInfo
    )

    data class UserAccountInfo(
        val username: String = "",
        val sessionId: String = ""
    )
}