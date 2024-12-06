import type { ComponentPropsWithoutRef } from 'react';

export function Button(props: ComponentPropsWithoutRef<'button'>) {
	return (
		<button
			{...props}
			className="bg-slate-300 text-black rounded text-xs px-2 leading-5 hover:brightness-95 active:brightness-90"
		/>
	);
}
