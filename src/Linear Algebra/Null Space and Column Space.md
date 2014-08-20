\title{Linear Algebra: Null Space and Column Space}
\date{August 18, 2014}
\author{Steven Schmatz, University of Michigan: College of Engineering}

\maketitle
\begin{center} stevenschmatz@gmail.com \end{center} \pagebreak

<!---Content goes here-->

# Matrix vector products

If you have an $m \times n$ matrix, there are simply $m$ rows and $n$ columns:

$$\textbf{A} =
\begin{bmatrix}
  a_{1,1} & a_{1,2} & \cdots & a_{1,n} \\
  a_{2,1} & a_{2,2} & \cdots & a_{2,n} \\
  \vdots  & \vdots  & \ddots & \vdots  \\
  a_{m,1} & a_{m,2} & \cdots & a_{m,n}
\end{bmatrix}$$

If the number of components in a vector is equal to the number of *columns* in a matrix, then you can multiply matrices and vectors.

$$
\textbf{A}\hat{x} = 
\begin{bmatrix}
a_{11}x_1 + a_{12}x_2 + \dots a_{1n}x_n \\
a_{21}x_1 + a_{22}x_2 + \dots a_{2n}x_n \\
\vdots \\
a_{m1}x_1 + a_{m2}x_2 + \dots a_{mn}x_n \\
\end{bmatrix} = \hat{b}
$$

It's essentially the dot product of each row vector in $\textbf{A}$ with the matrix $\hat{x}$! For example,

$$\begin{bmatrix}
-3 & 0 & 3 & 2 \\
1 & 7 & -1 & 9 \\
\end{bmatrix}
\begin{bmatrix}
2 \\ -3 \\ 4 \\ -1
\end{bmatrix} = 
\begin{bmatrix}
-6 + 0 + 12 - 2 \\
2 - 21 - 4 - 9
\end{bmatrix} =
\begin{bmatrix}
4 \\ -32
\end{bmatrix}
$$

### Transpose

The *transpose* function turns the columns of a vector into rows, and the rows of a vector into columns. This allows you to multiply two vectors. Hence, you can write the dot product of two vectors as the following:

$$\hat{a} \cdot \hat{b} = \hat{a}^\text{T}\hat{b}$$

# Introduction to the null space of matrix

If you take the set of all $\{\hat{x} \in \mathbb{R}^n \mid \textbf{A}\hat{x} = \textbf{0}\}$, do you have a valid subspace?

* Does it satisfy $\textbf{A} \hat{0} \in S$? Yes.

$$
\begin{bmatrix}
  a_{1,1} & a_{1,2} & \cdots & a_{1,n} \\
  a_{2,1} & a_{2,2} & \cdots & a_{2,n} \\
  \vdots  & \vdots  & \ddots & \vdots  \\
  a_{m,1} & a_{m,2} & \cdots & a_{m,n}
\end{bmatrix}
* 
\begin{bmatrix}
0 \\
0 \\
\vdots \\
0
\end{bmatrix} =
\begin{bmatrix}
0 \\
0 \\
\vdots \\
0
\end{bmatrix}
$$

* Does it have closure under addition? Yes.

$$\textbf{A}\hat{v_1} = \textbf{0}$$
$$\textbf{A}\hat{v_2} = \textbf{0}$$
$$\textbf{A}(\hat{v_1} + \hat{v_2}) = \textbf{0}$$
$$\hat{v_1} + \hat{v_2} \in N$$

* Does it have closure under scalar multiplication? Yes.

$$\hat{v_1} \in N$$
$$c\hat{v_1} \in N$$
$$\textbf{A}(c\hat{v_1}) = c\textbf{A} \hat{v_1} = c\textbf{0} = \textbf{0}$$

Hence, $N$ is a valid subspace, and is known as the *null space* of $A$.

# Calculating the null space of a matrix

$$\textbf{A} = \begin{bmatrix}
1 & 1 & 1 & 1 \\
1 & 2 & 3 & 4 \\
4 & 3 & 2 & 1
\end{bmatrix}
\begin{bmatrix}
x_1 \\ x_2 \\ x_3 \\ x_4
\end{bmatrix} = \begin{bmatrix} 0 \\ 0 \\ 0 \end{bmatrix}
$$
$$N(\textbf{A}) = \{\hat{x} \in \mathbb{R}^4 \mid \textbf{A}\hat{x} = \textbf{0}\}$$

Hence,

$$x_1 + x_2 + x_3 + x_4 = 0$$
$$x_1 + 2x_2 + 3x_3 + 4x_4 = 0$$
$$4x_1 + 3x_2 + 2x_3 + x_4 = 0$$

If we find the solution set to this, we have found our null space! We can represent this with an augmented matrix and put it in reduced row echelon form.

$$\begin{bmatrix}
1 & 1 & 1 & 1 & 0 \\
1 & 2 & 3 & 4 & 0 \\
4 & 3 & 2 & 1 & 0
\end{bmatrix}
$$

The zero column can actually never change, because whenever we do additions/multiplications, that row cannot change. So we are really just finding the reduced row echelon form of $\textbf{A}$!

$$\begin{bmatrix}
1 & 1 & 1 & 1 & 0 \\
0 & 1 & 2 & 3 & 0 \\
0 & 1 & 2 & 3 & 0
\end{bmatrix}
\rightarrow \begin{bmatrix}
1 & 0 & -1 & -2 & 0 \\
0 & 1 & 2 & 3 & 0 \\
0 & 0 & 0 & 0 & 0
\end{bmatrix}$$

Hence, this can be written as a system of equations as:

$$x_1 - x_3 - 2x_4 = 0$$
$$x_2 + 2x_3 + 3x_4 = 0$$

Or,

$$x_1 = x_3 + 2x_4$$
$$x_2 = -2x_3 - 3x_4$$

So you could write

$$\begin{bmatrix}
x_1 \\ x_2 \\ x_3 \\ x_4
\end{bmatrix} = x_3\begin{bmatrix}1 \\ -2 \\ 1 \\ 0\end{bmatrix} + x_4\begin{bmatrix}2\\-3\\0\\1\end{bmatrix}$$

You could pick any real number for $x_3$ and any real number for $x_4$! It's just a linear combination of those two vectors. Hence, one could write:

$$N(\textbf{A}) = \text{span}(\begin{bmatrix}1 \\ -2 \\ 1 \\ 0\end{bmatrix} \begin{bmatrix}2\\-3\\0\\1\end{bmatrix})$$

Since the finding the null space of the reduced row echelon form of $\textbf{A}$ is the same problem as finding the null space of $\textbf{A}$, one could write:

$$N(\textbf{A}) = N(\text{rref}(\textbf{A}))$$

# Relation of the relation of the null space to linear independence

If you're looking at the solution set:

$$\textbf{A}= \begin{bmatrix}\hat{v_1} + \hat{v_2} + \dots \hat{v_n}\end{bmatrix}\begin{bmatrix}x_1\\x_2\\ \vdots \\ x_n\end{bmatrix} = x_1\hat{v_1} + x_2\hat{v_2} + \dots x_n\hat{v_n} = \hat{0}$$

Then if all of your column vectors are *linearly independent*, that means that the *only* solution to $\textbf{A}\hat{x} = \hat{0}$ is $\hat{x} = \hat{0}$! $N(\textbf{A}) = \{\textbf{0}\}$. In addition, this means that $N(\textbf{A}) = N(\text{rref}(\textbf{A})) = \{0\}$, meaning no free variables.

$$N(\textbf{A}) = \{\textbf{0}\} \Longleftrightarrow \text{ column vectors of A are linearly independent.}$$

# Column space of a matrix

A *column space* of some matrix $\textbf{A} \in \mathbb{R}^{m \times n}$ with column vectors $\hat{v_1} \dots \hat{v_n}$ is the set of all linear combinations of the vectors, $C(\textbf{A}) = \text{span}(\hat{v_1} \dots \hat{v_n})$, and it is a valid subspace.

The column space is also defined as follows:

$$\{\textbf{A}\hat{x} \mid \hat{x} \in \mathbb{R}^n\}$$
$$= x_1\hat{v_1} + x_2\hat{v_2} + \dots x_n\hat{v_n} = \{x_1\hat{v_1} + x_2\hat{v_2} + \dots x_n\hat{v_n} \mid x_i \in \mathbb{R}\} = \text{span}(\hat{v_1}, \hat{v_2}, \dots \hat{v_n}) = C(\textbf{A})$$

If you needed to solve the equation:

$$\textbf{A}\hat{x} = \hat{b_1}$$

If $\hat{b_1} \notin C(\textbf{A}) \Longrightarrow \textbf{A}\hat{x} = \hat{b_1}$ has no solution! If $\textbf{A}\hat{x} = \hat{b_2}$ has at least one solution, that means $\hat{b_2} \in C(\textbf{A})$!

# Null space and column space basis

Say we have a matrix $\textbf{A}$:

$$\textbf{A} =
\begin{bmatrix}
1 & 1 & 1 & 1 \\
2 & 1 & 4 & 3 \\
3 & 4 & 1 & 2
\end{bmatrix}
$$

Then,

$$
C(\textbf{A}) = \text{span}(v_1 \dots v_4)
$$
$$N(\textbf{A}) = N(\text{rref}(\textbf{A}))$$
$$\text{rref}(\textbf{A}) =
\begin{bmatrix}
1 & 0 & 3 & 2 \\
0 & 1 & -2 & -1 \\
0 & 0 & 0 & 0
\end{bmatrix}
$$

So then what is the null space of this matrix?

$$
\begin{bmatrix}
1 & 0 & 3 & 2 \\
0 & 1 & -2 & -1 \\
0 & 0 & 0 & 0
\end{bmatrix}
\begin{bmatrix}
x_1 \\ x_2 \\ x_3 \\ x_4
\end{bmatrix}
= 
\begin{bmatrix}
0 \\ 0 \\ 0
\end{bmatrix}
$$
$$x_1 + 3x_3 + 2x_4 = 0$$
$$x_2 - 2x_3 - x_4 = 0$$

Solving for their pivot entries:

$$x_1 = -3x_3 - 2x_4$$
$$x_2 = 2x_3 + x_4$$

And hence, our solution set is:

$$N(\textbf{A}) = N(\text{rref}(\textbf{A})) =
\begin{bmatrix}
x_1 \\ x_2 \\ x_3 \\ x_4
\end{bmatrix}
=
x_3
\begin{bmatrix}
-3 \\ 2 \\ 1 \\ 0
\end{bmatrix}
+
x_4
\begin{bmatrix}
-2 \\ 1 \\ 0 \\ 1
\end{bmatrix}
$$

Hence,

$$N(\textbf{A}) = \text{span}(
\begin{bmatrix}
-3 \\ 2 \\ 1 \\ 0
\end{bmatrix},
\begin{bmatrix}
-2 \\ 1 \\ 0 \\ 1
\end{bmatrix})$$

Since the null space of $\textbf{A}$ has more than one solution, the column vectors of $\textbf{A}$ are linearly dependent! This also means that the column vectors are not a basis for the column space of $\textbf{A}$ (they are not linearly independent).

To find the basis of a column space, you simply eliminate the free variables. These are redundnant, so you can write:

$$C(\textbf{A}) = \text{span}(\hat{v_1}, \hat{v_2})$$

# Visualizing a column space as a plane in $\mathbb{R}^3$

Once you have found the basis for a column space, such as $C(\textbf{A}) = \text{span}(\hat{v_1}, \hat{v_2})$ in the last example, you can find the equation for a plane. Recall,

$$\hat{n} \cdot (\hat{x} - \begin{bmatrix}1\\2\\3\end{bmatrix}) = 0$$

If you take the cross product of the two vectors, you will get a vector that is normal to both vectors.

$$\hat{n} = \begin{bmatrix}1\\2\\3\end{bmatrix} \times \begin{bmatrix}1\\1\\4\end{bmatrix} = 
\begin{bmatrix}
8 - 3 \\
3 - 4 \\
1 - 2 
\end{bmatrix} =
\begin{bmatrix}
5 \\
-1 \\
-1
\end{bmatrix}
$$

Hence, one can write the equation of the plane:

$$\begin{bmatrix}5\\-1\\-1\end{bmatrix} \cdot (\begin{bmatrix}x\\y\\z\end{bmatrix} - \begin{bmatrix}1\\2\\3\end{bmatrix}) = 0$$
$$5(x-1) - 1(y-z) -1(z-3) = 0$$
$$5x-y-z=0$$

This plane intersects the origin! This makes sense because a column space *has to be a subspace*, which means it has to contain $\textbf{0}$. Another equivalent way to write $5x-y-z=0$

$$\{\textbf{A}\hat{x} \mid \hat{x} \in \mathbb{R}^n\}$$

Or, equivalently,

$$\{\hat{b} \mid \textbf{A}\hat{x} = \hat{b}, \hat{x}\in \mathbb{R}^n\}$$

Another way you could derive the solution set would be creating an augmented matrix of $\textbf{A}$ with $\hat{b} = (x, y, z)$, and finding the solutions using Gaussian elimination!

# Proof: Any subspace basis has the same number of elements

If you have a set $\textbf{A} \{a_1, a_2, \dots v_n\}$ which is a basis for $V$, then *any* spanning set **must** have a cardinality of $n$!

Let's say that you have a set $B = \{b_1, b_2, \dots b_n\}, m < n$ that spans $V$. If you construct a set $B_1' = \{a_1, b_1, b_2, \dots b_m\}$, this set is linearly dependent because at least one member of the set can be represented by a linear combination of the others: if $a_1 \in V$, and $\text{span}(b_1, b_2, \dots b_m) = V$, then the additional vector $a_1$ makes the set linearly dependent. Now, if you remove a single $b$ element from that list, you still have a set that spans $V$. 

If you substitute $b_2$ with $a_2$, then THAT set spans $V$. If you keep doing this process over and over again, you will essentially replace all of the $b$ terms with an $a$ term! However, if you have $m < n$, then $B_m$ with all elements replaced would span $V$. How can this work with a subset of $A$, which *itself* is linearly independent and spans $V$? It would mean that the original set $A$ is *linearly dependent*!

There cannot be a spanning set $B$ that has fewer elements than $A$.

Now we can define the *dimension* or *cardinality* of some subspace $V$ as the number of elements of any basis of $V$.

# Dimension of the null space or nullity

If we have some matrix $B$:

$$
B = \begin{bmatrix}
1 & 1 & 2 & 3 & 2 \\
1 & 1 & 3 & 1 & 4
\end{bmatrix}, \text{rref}(\textbf{B}) =
\begin{bmatrix}
1 & 1 & 0 & 7 & -2 \\
0 & 0 & 1 & -2 & 2
\end{bmatrix}
$$

Hence,

$$x_1 + x_2 + 7x_4 -2x_5 = 0$$
$$x_3 - 2x_4 + 2x_5 = 0$$
$$x_1 = -x_2 - 7x_4 + 2x_5$$
$$x_3 = 2x_4 - 2x_5$$

And:

$$\begin{bmatrix}x_1 \\ x_2 \\ x_3 \\ x_4 \\ x_5\end{bmatrix} =
x_2 \begin{bmatrix}-1 \\ 1 \\ 0 \\ 0 \\ 0 \end{bmatrix} + 
x_4 \begin{bmatrix}-7 \\ 0 \\ 2 \\ 1 \\ 0\end{bmatrix} + 
x_5 \begin{bmatrix}2 \\ 0 \\ -2 \\0 \\1 \end{bmatrix}
$$
$$N(\textbf{A}) = N(\text{rref}(\textbf{A}))$$

Can one of the vectors be represented as a linear combination of the other vectors? **No**, because each vector has a one where the other vectors have a zero. Hence,

$$\{\hat{v_1}, \hat{v_2}, \hat{v_3}\} \text{ is a basis for } N(\textbf{B})$$
$$\text{dim}(N(\textbf{B})) = 3 = \text{nullity}(\textbf{B})$$

The nullity of B is the number of free variables that you have, because each has a linearly independent vector.

# Dimension of the column space or rank

How do you find a basis for the column space of a matrix? For example,

$$
\textbf{A} = \begin{bmatrix}
1 & 0 & -1 & 0 & 4 \\
2 & 1 & 0 & 0 & 9 \\
-1 & 2 & 5 & 1 & -5 \\
1 & -1 & -3 & -2 & 9
\end{bmatrix}
$$

1. Put the matrix in reduced row echelon form

$$
\text{rref}(\textbf{A}) = \begin{bmatrix}
1 & 0 & -1 & 0 & 4 \\
0 & 1 & 2 & 0 & 1 \\
0 & 0 & 0 & 1 & -3 \\
0 & 0 & 0 & -2 & 6
\end{bmatrix}
$$

2. Find all free variable vectors (those are linearly dependent). The corresponding columns in $A$ are *also* linearly independent! 

It turns out that you can represent non-pivot columns as linear combinations of the pivot columns! Hence, the dimension of the column space, or the *rank*, is 3.

# Showing relation between basis cols and pivot cols

So, why did the last example work? By definition, all of the pivot cols in the reduced row echelon form of a matrix are linearly independent. We know that $N(\text{rref}(\textbf{A})) = N(\textbf{A})$, so once we find the pivot cols then those are the linearly independent basis columns.

This also must mean that the basis must span C(A). If you can show that your free variable columns can be represented as a linear combination of the other two columns, then you can say that the pivot columns do indeed span $C(\textbf{A})$.

To show this, consider that you can set the value of the free variables to any value. This means that you could set the linear combination of all columns to the zero vector, subtract one free variable to the other side, and flip its sign. Hence, you could show that the free variable columns can be rewritten as a linear combination of the pivot columns.