// src/pages/HomePage.js
import React from 'react'
import EventList from '../components/EventList/EventList'
import EventForm from '../components/EventForm/EventForm'
import '../styles/global.css' 

const HomePage = () => {
	return (
		<div className='home-page'>
			<EventForm /> 
			<EventList /> 
		</div>
	)
}

export default HomePage
