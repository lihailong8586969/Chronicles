List<String> strings = new ArrayList<>();
strings.add("Hello");
strings.add("World");
// strings.add(new Date());    #1
// Integer i = strings.get(0); #1
for (String s : strings) {   //#2
 System.out.printf("%s has length %d%n", s, s.length());
}
// 1 - won't comile
// 2 - for-each loop knows the contained data type is String

public interface List<E> extends Collection<E> // E is the type parameter
boolean add(E e) // 1
boolean addAll(Collection<? extends E> c) // 2
void clear() // 3
boolean contains(Object o) // 3
boolean containsAll(Collection<?> c) // 4
E get(int index) // 1
 
// 1 - Using the type parameter E as an argument or return type
// 2 - A bounded wildcard
// 3 - Methods that do not involve the type itself
// 4 - An unknown type

static <T> List<T> emptyList()
static <K,V> Map<K,V> emptyMap()
static <T> boolean addAll(Collection<? super T> c, T... elements)
static <T extends Object & Comparable<? super T>> T min(Collection<? extends T> coll)
 
// below code is Ok 
List<Object> objects = new ArrayList<Object>();
objects.add("Hello");
objects.add(LocalDate.now());
objects.add(3);
System.out.println(objects);

List<String> strings = new ArrayList<>();
String s = "abc";
Object o = s;  // 1
// strings.add(o); // 2
// List<Object> moreObjects = strings; // 3
// moreObjects.add(new Date());
// String s = moreObjects.get(0); // 4

// 1 - Allowed
// 2 - Not allowed
// 3 - Also not allowed, but pretend it was
// 4 - Corrupted collection
// The problem is that List<String> is not a subclass of List<Object>
// We say that the parameterized type is invariant.


// Unbounded Wildcards
// If you declare a List of unbounded type, you can read from it but not write to it.
List<?> stuff = new ArrayList<>();
// stuff.add("abc");  // 1
// stuff.add(new Object());
// stuff.add(3);
int numElements = stuff.size(); // 2
// 1 - No additions allowed
// 2 - numElements is zero

// One use for them is that any method that takes a List<?> as an argument will accept any list at all when invoked

// Upper Bounded Wildcards
List<? extends Number> numbers = new ArrayList<>();
// numbers.add(3); // Still cannot add values
// numbers.add(3.14159);
// numbers.add(new BigDecimal("3"));
// The problem is that when you retrieve the value, the compiler has no idea what type it is, only that it extends Number

private static double sumList(List<? extends Number> list) {
 Number num = list.get(0);
 // ... from before ...
}
// When the method is invoked, the elements of the list will be either Number or one of its descendants, 
//   so a Number reference will always be correct.
