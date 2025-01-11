// src/App.js
import React from 'react'
import { Route, Routes, Navigate } from 'react-router-dom'
import HomePage from './pages/HomePage'
import Auth from './pages/Auth'

const App = () => {
	return (
		<Routes>
			<Route path='/auth' element={<Auth />} />
			<Route path='/home' element={<HomePage />} />
			<Route path='*' element={<Navigate to='/auth' />} />
		</Routes>
	)
}

export default App
