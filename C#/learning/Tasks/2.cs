using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Threading;

namespace FlyBird
{
    public partial class Form2 : Form
    {

        private Form form = null;
        private PictureBox bird = null;
        private Task task;
        private CancellationTokenSource cts = new CancellationTokenSource();

        public Form2()
        {
            InitializeComponent();
        }

        public void Form2_Load(object sender, EventArgs e)
        {

            this.form = sender as Form;
            this.bird = this.pictureBox1;

            System.Windows.Forms.Control.CheckForIllegalCrossThreadCalls = false;
            this.task = new Task(() => {

                while (true){

                    if (this.cts.Token.IsCancellationRequested)
                    {
                        break;
                    }
                    Thread.Sleep(900);
                    this.fly();
                }
            });
            task.Start();
        }

        public void fly(string direction = "down")
        {

            this.bird.Top += (direction == "down") ? 5 : -5;
        }

        public void onKeyDown(object sender, KeyEventArgs e)
        {
            this.cts.Cancel();
            if (e.KeyCode == Keys.Up){

                this.fly("up");
                this.task = new Task(() => {

                    this.cts = new CancellationTokenSource();
                    while (true)
                    {

                        if (this.cts.Token.IsCancellationRequested)
                        {
                            break;
                        }
                        Thread.Sleep(900);
                        this.fly();
                    }
                });
                task.Start();
                return;
            }

            this.fly("down");
            this.task = new Task(() => {

                this.cts = new CancellationTokenSource();
                while (true)
                {

                    if (this.cts.Token.IsCancellationRequested)
                    {
                        break;
                    }
                    Thread.Sleep(900);
                    this.fly();
                }
            });
            task.Start();
        }

    }
}
