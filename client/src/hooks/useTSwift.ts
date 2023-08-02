import { useContext } from "react"
import { DocumentFile, TSwiftContext } from "../context"
import { fireDangerToast, fireScucessToast } from "../components/toasts"

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

    const runProgram = () => {
        dispatch({ type: 'set-terminal-content', payload: { content: '' } })
        dispatch({ type: 'reset-graphviz-content' })
        dispatch({ type: 'set-errors', payload: { errors: [] } })
        setSymbolTable('')


        const programInput = state.currentDocument.content
        setTerminalContent('Ejecutando programa...')

        // Todo: fetch TSwift runtime


        // * Set terminal content
        // setTerminalContent()

        // * Set errors
        // dispatch({ type: 'set-errors', payload: { errors: runtime.ast.context.errorTable.errors } })

        // * Fire toast
        // if (runtime.ast.context.errorTable.errors.length > 0) {
        //     fireDangerToast('Programa ejecutado con errores')
        // } else {
        //     fireScucessToast('Programa ejecutado con éxito')
        // }

        // * CST graphviz report
        // dispatch({ type: 'set-graphviz-content', payload: { content: runtime.ast.graphviz } })

        // * Symbol table report
        // setSymbolTable(runtime.ast.context.scopeTrace.graphviz || '');

        // * Reset terminal content
        dispatch({ type: 'reset-graphviz-content' })
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
