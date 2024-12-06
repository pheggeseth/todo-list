import type { ReactNode } from 'react';
import '../styles.css';

type RootLayoutProps = { children: ReactNode };

export default async function RootLayout({ children }: RootLayoutProps) {
	const data = await getData();

	return (
		<div className="font-['Nunito']">
			<meta name="description" content={data.description} />
			<header>
				<h1>Go Todos</h1>
			</header>
			<main className="m-6 flex items-center *:min-h-64 *:min-w-64 lg:m-0 lg:min-h-svh lg:justify-center">
				{children}
			</main>
		</div>
	);
}

const getData = async () => {
	const data = {
		description: 'An internet website!',
		icon: '/images/favicon.png',
	};

	return data;
};

export const getConfig = async () => {
	return {
		render: 'static',
	} as const;
};
