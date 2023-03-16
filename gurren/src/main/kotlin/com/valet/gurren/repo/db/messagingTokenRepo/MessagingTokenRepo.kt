package com.valet.gurren.repo.db.messagingTokenRepo

import com.valet.gurren.model.MessagingToken
import org.springframework.data.repository.CrudRepository

interface MessagingTokenRepo : CrudRepository<MessagingToken, String> {
}