// src/pages/Auth.js
import React, { useState } from 'react'
import axios from 'axios'
import '../styles/global.css'
import { Link, useNavigate } from 'react-router-dom'

const Auth = () => {
	const [isLogin, setIsLogin] = useState(true) // Состояние для переключения между входом и регистрацией
	const [login, setLogin] = useState('') // Логин
	const [email, setEmail] = useState('') // Email для регистрации
	const [password, setPassword] = useState('')
  const navigate = useNavigate()

	const handleSubmit = async e => {
		e.preventDefault()
		const url = isLogin
			? 'http://localhost:8000/auth'
			: 'http://localhost:8000/register'

		try {
			const payload = isLogin ? { login, password } : { login, email, password } // Добавлено поле email для регистрации
			const response = await axios.post(url, payload)
			console.log(response.data)

      if(response.status === 200){
        navigate('/home')
      }

		} catch (error) {
			console.error('Ошибка при авторизации:', error)
		}
	}

	return (
		<div className='auth-container'>
			<h2>{isLogin ? 'Authorization' : 'Registration'}</h2>
			<form id='form-auth' className={isLogin ? 'show' : 'hide'} onSubmit={handleSubmit}>
				<input
					type='text'
					placeholder='Login'
					value={login}
					onChange={e => setLogin(e.target.value)}
					required
				/>
				{!isLogin && ( // Поле для email отображается только при регистрации
					<input
						type='email'
						placeholder='Email'
						value={email}
						onChange={e => setEmail(e.target.value)}
						required
					/>
				)}
				<input
					type='password'
					placeholder='Password'
					value={password}
					onChange={e => setPassword(e.target.value)}
					required
				/>
				<button type='submit'>{isLogin ? 'Auth' : 'Register'}</button>
			</form>
			<p onClick={() => setIsLogin(!isLogin)}>
				{isLogin
					? 'Dont have an account? Sign up'
					: 'Already have an account? Sign in'}
			</p>
		</div>
	)
}

export default Auth
