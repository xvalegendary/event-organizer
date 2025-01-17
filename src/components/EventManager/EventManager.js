// src/components/EventManager/EventManager.js
import React, { useState } from 'react'
import EventForm from '../EventForm/EventForm'
import EventList from '../EventList/EventList'

const EventManager = () => {
	const [events, setEvents] = useState([])

	const handleEventCreate = newEvent => {
		setEvents([...events, newEvent]) // Добавляем новое событие в список
	}

	return (
		<div>
			<EventList events={events} />
		</div>
	)
}

export default EventManager
