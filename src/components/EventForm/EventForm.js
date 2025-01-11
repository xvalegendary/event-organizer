// src/components/EventForm/EventForm.js
import React, { useState } from 'react'
import './EventForm.css'

const EventForm = () => {
	const [eventName, setEventName] = useState('')
	const [eventDate, setEventDate] = useState('')
	const [eventLocation, setEventLocation] = useState('')
	const [eventTime, setEventTime] = useState('')

	const handleSubmit = e => {
		e.preventDefault()

		console.log('Created an event:', {
			eventName,
			eventDate,
			eventTime,
			eventLocation,
		})

		// Сброс полей формы
		setEventName('')
		setEventDate('')
		setEventTime('')
		setEventLocation('')
	}

	return (
		<div className='event-form-container'>
			<form className='event-form' onSubmit={handleSubmit}>
				<h2>Create an Event:</h2>
				<input
					type='text'
					id='input-1'
					placeholder='Event Name'
					value={eventName}
					onChange={e => setEventName(e.target.value)}
					required
				/>
				<input
					type='date'
					id='input-2'
					value={eventDate}
					onChange={e => setEventDate(e.target.value)}
					required
				/>
				<input
					type='time'
					id='input-3'
					value={eventTime}
					onChange={e => setEventTime(e.target.value)}
					required
				/>
				<input
					type='text'
					id='input-4'
					placeholder='Location of the Event'
					value={eventLocation}
					onChange={e => setEventLocation(e.target.value)}
					required
				/>
				<button type='submit'>Create Event</button>
			</form>
		</div>
	)
}

export default EventForm
