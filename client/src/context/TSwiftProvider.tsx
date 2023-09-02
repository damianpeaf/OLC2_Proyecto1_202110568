import { FC, ReactNode, useReducer, useEffect } from 'react';
import { TSwiftActionType, TSwiftContext, TSwiftReducer } from '.';
import { SymbolTableI } from '../components/modal';

export interface DocumentFile {
    id: number;
    name: string;
    content: string;
}

export const initialDocument: DocumentFile = {
    id: 0,
    name: 'Untitled',
    content: ''
}

export interface TSwiftError {

    Column: number;
    Line: number;
    Msg: string;
    Type: string;

}

export interface TSwiftState {
    isConsoleOpen: boolean;
    documents: DocumentFile[];
    currentDocument: DocumentFile;
    isRenameModalOpen: boolean;
    terminalContent: string;
    isAstModalOpen: boolean;
    graphviz: string | null;
    errors: TSwiftError[];
    symbolTable: SymbolTableI | null;
    isSymbolTableModalOpen: boolean;
}

interface TSwiftProviderProps {
    children: ReactNode
}

const defaultState: TSwiftState = {
    isConsoleOpen: false,
    documents: [initialDocument],
    currentDocument: initialDocument,
    isRenameModalOpen: false,
    terminalContent: '',
    isAstModalOpen: false,
    graphviz: null,
    errors: [],
    symbolTable: null,
    isSymbolTableModalOpen: false
}

const TSwift_INITIAL_STATE: TSwiftState =
    localStorage.getItem('state')
        ? JSON.parse(localStorage.getItem('state')!)
        : defaultState

export const TSwiftProvider: FC<TSwiftProviderProps> = ({ children }) => {

    const [state, dispatch] = useReducer(TSwiftReducer, TSwift_INITIAL_STATE)

    const saveState = () => {
        localStorage.setItem('state', JSON.stringify(state))
    }

    useEffect(() => {
        saveState()
    }, [state])

    return (
        < TSwiftContext.Provider value={{
            ...state,
            dispatch,
            saveState
        }}>
            {children}
        </ TSwiftContext.Provider>
    )
}