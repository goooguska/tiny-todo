export type Task = {
	id: string;
	title: string;
	description: string;
	completed: boolean;
};

export type CreateTaskInput = {
	title: string;
	description: string;
	completed?: boolean;
};

