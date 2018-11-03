


Please! please! please! DO NOT serialize data and place it into your database. Serialize can be used that way, but that's missing the point of a relational database and the datatypes inherent in your database engine. Doing this makes data in your database non-portable, difficult to read, and can complicate queries. If you want your application to be portable to other languages, like let's say you find that you want to use Java for some portion of your app that it makes sense to use Java in, serialization will become a pain in the buttocks. You should always be able to query and modify data in the database without using a third party intermediary tool to manipulate data to be inserted. 

I've encountered this too many times in my career, it makes for difficult to maintain code, code with portability issues, and data that is it more difficult to migrate to other RDMS systems, new schema, etc. It also has the added disadvantage of making it messy to search your database based on one of the fields that you've serialized. 

That's not to say serialize() is useless. It's not... A good place to use it may be a cache file that contains the result of a data intensive operation, for instance. There are tons of others... Just don't abuse serialize because the next guy who comes along will have a maintenance or migration nightmare.


注意！注意！ 请勿序列化数据并将其放入数据库。 Serialize可以这种方式使用，但是缺少关系数据库和数据库引擎中固有的数据类型。这样做会使数据库中的数据不可移植，难以阅读，并且可能使查询复杂化。如果你希望你的应用程序可以移植到其他语言，比如让你说你发现你想在你的应用程序的某些部分使用Java，那么使用Java是有意义的，序列化将成为臀部的痛苦。您应始终能够查询和修改数据库中的数据，而无需使用第三方中间工具来操作要插入的数据。

我在职业生涯中遇到过这么多次，这使得难以维护代码，代码存在可移植性问题，而且数据更难以迁移到其他RDMS系统，新模式等。它还有另外的缺点根据您已序列化的某个字段搜索数据库使其变得混乱。

这并不是说serialize（）没用。它不是......使用它的好地方可能是包含数据密集型操作结果的缓存文件。还有很多其他的......只是不要滥用序列化，因为下一个出现的人会有维护或迁移的噩梦。