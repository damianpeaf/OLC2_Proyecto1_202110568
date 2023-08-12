import { MdOutlineCloseFullscreen } from 'react-icons/md';
import { useTSwift } from '../../hooks';
import React, { useState, useEffect } from 'react';
import { ErrorTable } from './ErrorTable';

export const Console = () => {

    const { isConsoleOpen, closeTerminal, terminalContent, errors } = useTSwift()

    return (
        <section
            className={`
            ${isConsoleOpen ? 'scale-100' : 'scale-0'}
            absolute
            bottom-0
            left-0
            w-full
            h-1/3
            bg-background-dark
            transition-all
            `}
        >
            <article
                className="
                flex
                justify-between
                px-4
                mt-2
                "
            >
                <h2
                    className="
                    text-gray-300
                    font-bold
                    text-xl
                    "
                >
                    Consola
                </h2>
                <button
                    className="
                        p-2
                        text-gray-300
                        font-bold
                        text-xl
                    "
                    onClick={closeTerminal}
                >
                    <MdOutlineCloseFullscreen />
                </button>
            </article>
            <article
                className="
                    py-2
                    px-4
                    overflow-y-auto
                    mb-2
                    h-4/5
                    console-font
                    text-gray-400
                "
            >


                {
                    errors.length > 0 &&
                    <>
                        <br />
                        <h3
                            className='text-gray-300 font-bold text-xl'
                        >Errores:</h3>
                        <br />
                        <ErrorTable errors={errors} />
                        <br />
                        <br />
                        <h3
                            className='text-gray-300 font-bold text-xl'
                        >Salida:</h3>
                        <br />
                    </>
                }

                {
                    terminalContent.split('\n').map((line, index) => (
                        <React.Fragment key={index}>
                            <pre className='whitespace-pre-wrap'>
                                {line.replace(/\t/g, '\u00a0\u00a0\u00a0\u00a0')}
                            </pre>
                            {index !== terminalContent.split('\n').length - 1 && <br />}
                        </React.Fragment>
                    ))
                }
            </article>
        </section>
    )
}
