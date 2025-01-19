import React from 'react'

import axios from "axios";

interface Todo {
  completed: boolean
  id: number
  text: string
}

function App() {
  const [currentTodos, setCurrentTodos] = React.useState<Todo[]>([])

  const todoInput = React.useRef<HTMLInputElement>(null)

  const addTodo = (text: string) => {
    setCurrentTodos((prev) => [
      ...prev,
      {
        completed: false,
        id: Math.random(),
        text,
      },
    ])
  }

  const toggleTodo = (id: number) => {
    setCurrentTodos((prev) =>
      prev.map((todo) =>
        todo.id === id ? { ...todo, completed: !todo.completed } : todo
      )
    )
  }

  const deleteTodo = (id: number) => {
    setCurrentTodos((prev) => prev.filter((todo) => todo.id !== id))
  }

  const deleteCompletedTodos = () => {
    setCurrentTodos((prev) => prev.filter((todo) => !todo.completed))
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-700">
      <div className="bg-gray-600 rounded-md shadow-md p-2">
        <h1 className="text-2xl">Todo app with a Go backend</h1>

        {currentTodos.map((todo) => {
          return (
            <div key={todo.id} className="flex items-center space-x-2 flex items-center">
              <div className={`${todo.completed ? "bg-blue-300 border-blue-300" : "border-gray-400"} rounded-full border border-gray-400 p-2 cursor-pointer`} onClick={() => toggleTodo(todo.id)}>

              </div>

              <span
                className={todo.completed ? 'line-through' : ''}
              >
                {todo.text}
              </span>
            </div>
          )
        })}

        <div className="flex flex-col w-full items-center mt-2">
          <input type="text" className="rounded-md border outline-none shadow-md p-1 text-white bg-gray-700 border-gray-700 w-full" ref={todoInput} />
          <button className="w-full bg-blue-400 text-white p-2 rounded-md shadow-md mt-2" onClick={() => addTodo(todoInput.current!.value)}>Add</button>

          <button className="w-full bg-blue-400 text-white p-2 rounded-md shadow-md mt-2" onClick={deleteCompletedTodos}>Delete completed</button>
        </div>
      </div>
    </div>
  )
}

export default App
