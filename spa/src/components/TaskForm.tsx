import { useRef } from 'react'

type Props = {
  addTask: (title: string) => void
}

export default function TaskForm({ addTask }: Props) {
  const inputEl = useRef<HTMLInputElement>(null)
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (inputEl.current) {
      addTask(inputEl.current.value)
      inputEl.current.value = ''
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" ref={inputEl} placeholder="New task title" />
      <button type="submit">Submit</button>
    </form>
  )
}