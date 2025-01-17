// src/pages/HomePage.js
import React from 'react'
import EventList from '../components/EventList/EventList'
import EventForm from '../components/EventForm/EventForm'
import EventManager from '../components/EventManager/EventManager'
import '../styles/global.css' 

const HomePage = () => {
	return (
		<div className='home-page'>
			<EventForm /> 
			<EventList /> 
			<EventManager />
		</div>
	)
}

export default HomePage
