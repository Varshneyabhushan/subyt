import React from "react";

interface ErrorBoundaryState {
  hasError : boolean;
  message : string;
}

//TODO implement showing custom message
export default class ErrorBoundary extends React.Component <any, ErrorBoundaryState>{
    constructor(props: any) {
      super(props);
      this.state = { hasError: false, message : "" };
    }
  
    static getDerivedStateFromError(error : Error) {
      // Update state so the next render will show the fallback UI.
      return { hasError: true };
    }
  
    componentDidCatch(error : Error, info : any) {
      // info.componentStack
      // Example "componentStack":
      //   in ComponentThatThrows (created by App)
      //   in ErrorBoundary (created by App)
      //   in div (created by App)
      //   in App
      console.log(error, info);
    }
  
    render() {
      if (this.state.hasError) {
        // You can render any custom fallback UI
        return this.props.fallback
      }
  
      return this.props.children;
    }
  }