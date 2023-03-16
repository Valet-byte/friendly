package com.valet.gurren.service.userService

import com.valet.gurren.model.Description
import com.valet.gurren.model.User
import com.valet.gurren.model.UserData
import com.valet.gurren.model.UserProfileSetting
import org.springframework.http.ResponseEntity
import org.springframework.stereotype.Service

@Service
class UserServiceImpl(

) : UserService {
    override fun login(userId: String, sessionId: String, userData: UserData): User {
        TODO("Not yet implemented")
    }

    override fun updateSessionId(sessionId: String, userId: String): ResponseEntity<String> {
        TODO("Not yet implemented")
    }

    override fun getUserFriends(userId: String): ResponseEntity<List<User>> {
        TODO("Not yet implemented")
    }

    override fun addFriend(userId: String, username: String): ResponseEntity<String> {
        TODO("Not yet implemented")
    }

    override fun getSubscriberList(userId: String): ResponseEntity<List<User>> {
        TODO("Not yet implemented")
    }

    override fun deleteFriend(userId: String, username: String): ResponseEntity<String> {
        TODO("Not yet implemented")
    }

    override fun deleteUser(userId: String): ResponseEntity<String> {
        TODO("Not yet implemented")
    }

    override fun getUserDescription(username: String): ResponseEntity<Description> {
        TODO("Not yet implemented")
    }

    override fun getUserInfo(username: String): ResponseEntity<User> {
        TODO("Not yet implemented")
    }

    override fun updateUserProfileSettings(
        userId: String,
        userProfileSetting: UserProfileSetting
    ): ResponseEntity<String> {
        TODO("Not yet implemented")
    }

    override fun getUserProfileSettings(userId: String): ResponseEntity<UserProfileSetting> {
        TODO("Not yet implemented")
    }
}