// src/components/EventForm/EventForm.js
import React, { useState } from 'react'
import './EventForm.css'

const EventForm = () => {
	const [eventName, setEventName] = useState('')
	const [eventDate, setEventDate] = useState('')
	const [eventLocation, setEventLocation] = useState('')
	const [eventTime, setEventTime] = useState('')
	const [events, setEvents] = useState([])

	const handleSubmit = e => {
		e.preventDefault()
		const newEvent = {
			eventName,
			eventDate,
			eventTime,
			eventLocation,
		}
		console.log('Created an event:', newEvent)
		setEvents([...events, newEvent])
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
			<div className='event-list-container'>
				<h3>List of Events:</h3>
				<ul className='event-list'>
					{events.length > 0 ? (
						events.map((event, index) => (
							<li key={index}>
								<p>
									<strong>Event Name:</strong> {event.eventName}
								</p>
								<p>
									<strong>Date:</strong> {event.eventDate}
								</p>
								<p>
									<strong>Time:</strong> {event.eventTime}
								</p>
								<p>
									<strong>Location:</strong> {event.eventLocation}
								</p>
							</li>
						))
					) : (
						<p>No events created yet.</p>
					)}
				</ul>
			</div>
		</div>
	)
}

export default EventForm
