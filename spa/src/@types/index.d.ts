declare type Task = {
  id: number,
  title: string,
  completed: boolean,
  createdAt: string,
  updatedAt: string
}

declare type TaskUpdateParams = {
  title?: string
}
