import React from "react";

interface ErrorBoundaryState {
  hasError : boolean;
}

//TODO implement showing custom message
export default class ErrorBoundary extends React.Component <any, ErrorBoundaryState>{
    constructor(props: any) {
      super(props);
      this.state = { hasError: false };
    }
  
    static getDerivedStateFromError(error : Error) {
      return { hasError: true };
    }
  
    componentDidCatch(error : Error, info : any) {
      console.log(error, info);
    }
  
    render() {
      if (this.state.hasError) {
        return this.props.fallback
      }
  
      return this.props.children;
    }
  }