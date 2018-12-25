System.Windows.Forms.Control.CheckForIllegalCrossThreadCalls = false;
this.thread = new Thread(new ThreadStart( ()=> { this.fly(); } ) );
this.thread.Start();