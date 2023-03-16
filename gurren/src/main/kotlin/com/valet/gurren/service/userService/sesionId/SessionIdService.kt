package com.valet.gurren.service.userService.sesionId

interface SessionIdService {
    fun updateSessionId(uid: String, username: String, sessionId: String)

}