Task.Factory.StartNew(delegate
{
    for(int i=1;i<50;i++)
    {
        System.Diagnostics.Debug.WriteLine(i.ToString());
        Thread.Sleep(100);
    }
});


Task.Factory.StartNew(delegate
{
    for (int i = 101; i < 150; i++)
    {
        System.Diagnostics.Debug.WriteLine(i.ToString());
        Task.Delay(100);
    }
});


