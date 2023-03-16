package com.valet.gurren.service.userService

import com.valet.gurren.model.Description
import com.valet.gurren.model.User
import com.valet.gurren.model.UserProfileSetting
import org.springframework.http.ResponseEntity

interface UserService {
    fun updateSessionId(sessionId: String, userId: String) : ResponseEntity<String>
    fun getUserFriends(userId: String) : ResponseEntity<List<User>>
    fun addFriend(userId: String, username: String) : ResponseEntity<String>
    fun getSubscriberList(userId: String) : ResponseEntity<List<User>>
    fun deleteFriend(userId: String, username: String) : ResponseEntity<String>
    fun deleteUser(userId: String) : ResponseEntity<String>
    fun getUserDescription(username: String) : ResponseEntity<Description>
    fun getUserInfo(username: String) : ResponseEntity<User>
    fun updateUserProfileSettings(userId: String, userProfileSetting: UserProfileSetting) : ResponseEntity<String>
    fun getUserProfileSettings(userId: String) : ResponseEntity<UserProfileSetting>
}