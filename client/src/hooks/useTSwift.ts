import { useContext } from "react"
import { DocumentFile, TSwiftContext, TSwiftError } from "../context"
import { fireDangerToast, fireScucessToast } from "../components/toasts"

export type ApiResponse = {
    output: string
    errors: TSwiftError[] | null
    cstSvg: string
}

export const useTSwift = () => {

    const { dispatch, ...state } = useContext(TSwiftContext)

    const openTerminal = () => {
        dispatch({ type: 'open-terminal' })
    }

    const closeTerminal = () => {
        dispatch({ type: 'close-terminal' })
    }


    const renameDocument = (id: number, name: string) => {
        dispatch({ type: 'rename-tab', payload: { id, name } })
    }

    const closeDocument = (id: number) => {
        dispatch({ type: 'close-tab', payload: { id } })
    }

    const openDocument = (document: Omit<DocumentFile, 'id'>) => {
        dispatch({ type: 'open-file', payload: { document } })
    }

    const newDocument = () => {
        dispatch({ type: 'new-tab' })
    }

    const saveDocument = (document: DocumentFile) => {
        dispatch({ type: 'save-file', payload: { document } })
    }

    const setCurrentDocument = (id: number) => {
        dispatch({ type: 'set-current-document', payload: { id } })
    }

    const openRenameModal = () => {
        dispatch({ type: 'open-rename-modal' })
    }

    const closeRenameModal = () => {
        dispatch({ type: 'close-rename-modal' })
    }

    const setTerminalContent = (content: string) => {
        dispatch({ type: 'set-terminal-content', payload: { content } })
    }

    const setSymbolTable = (content: string) => {
        dispatch({ type: 'set-symbol-table', payload: { content } })
    }

    const runProgram = async () => {
        dispatch({ type: 'set-terminal-content', payload: { content: '' } })
        dispatch({ type: 'reset-graphviz-content' })
        dispatch({ type: 'set-errors', payload: { errors: [] } })
        setSymbolTable('')


        const programInput = state.currentDocument.content
        setTerminalContent('Ejecutando programa...')

        // form-data
        const formData = new FormData()
        formData.append('code', programInput)

        const res = await fetch(import.meta.env.VITE_API_URL + '/compile', {
            method: 'POST',
            body: formData
        })

        const { errors, output, cstSvg } = await res.json() as ApiResponse


        if (cstSvg != null) {
            console.log({
                type: 'set-graphviz-content',
                payload: { content: cstSvg }
            })
            dispatch({ type: 'set-graphviz-content', payload: { content: cstSvg } })
        }

        // * Set terminal content
        setTerminalContent(output)

        // * Set errors

        // * Fire toast
        if (errors) {
            dispatch({ type: 'set-errors', payload: { errors } })
            fireDangerToast('Programa ejecutado con errores')
        } else {
            fireScucessToast('Programa ejecutado con éxito')
        }

        // * CST graphviz report

        // * Symbol table report
        // setSymbolTable(runtime.ast.context.scopeTrace.graphviz || '');

        // * Reset terminal content
        // dispatch({ type: 'reset-graphviz-content' })
        setSymbolTable('')

    }

    const openAstModal = () => {
        dispatch({ type: 'open-ast-modal' })
    }

    const closeAstModal = () => {
        dispatch({ type: 'close-ast-modal' })
    }

    const toogleTerminal = () => {
        if (state.isConsoleOpen) {
            closeTerminal()
        } else {
            openTerminal()
        }
    }

    const closeSymbolTableModal = () => {
        dispatch({ type: 'close-symbol-table-modal' })
    }

    const openSymbolTableModal = () => {
        dispatch({ type: 'open-symbol-table-modal' })
    }

    return {
        ...state,
        openTerminal,
        closeTerminal,
        renameDocument,
        closeDocument,
        openDocument,
        newDocument,
        saveDocument,
        setCurrentDocument,
        openRenameModal,
        closeRenameModal,
        setTerminalContent,
        runProgram,
        openAstModal,
        closeAstModal,
        toogleTerminal,
        setSymbolTable,
        closeSymbolTableModal,
        openSymbolTableModal
    }
}
