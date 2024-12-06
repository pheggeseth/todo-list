export default async function HomePage() {
	const todos = await fetch('http://localhost:8080/api/todos').then((res) =>
		res.json(),
	);

	return (
		<div>
			<title>Go Todos</title>
			<pre>{JSON.stringify(todos, null, 2)}</pre>
		</div>
	);
}

export const getConfig = async () => {
	return {
		render: 'static',
	} as const;
};
