import { ExecutionState } from './ExecutionState';

export type ExecutionCallback = (status: ExecutionState) => void;