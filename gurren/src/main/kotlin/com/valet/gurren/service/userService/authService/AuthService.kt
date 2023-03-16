package com.valet.gurren.service.userService.authService

import com.valet.gurren.model.User
import com.valet.gurren.model.UserData

interface AuthService {
    fun auth(userId: String): User
    fun registration(userId: String, userData: UserData)
}