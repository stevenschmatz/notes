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

It's essentially the dot product of each row vector with the matrix $\hat{x}$! For example,

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

The *transpose* function turns the columns of a vector into rows, and the rows of a vector into columns. This allows you to multiply two vectors.