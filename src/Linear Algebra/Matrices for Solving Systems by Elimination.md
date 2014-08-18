\title{Linear Algebra: Matrices for Solving Systems by Elimination}
\date{August 18, 2014}
\author{Steven Schmatz, University of Michigan: College of Engineering}

\maketitle
\begin{center} stevenschmatz@gmail.com \end{center} \pagebreak

<!---Content goes here-->

# Matrices: Reduced row echelon form

You can construct a coefficient matrix from a system of linear equations, with each element in the matrix corresponding to a coefficient in the system of linear equations. For example, a valid system of linear equations would be:

$$x_1 + 2x_2 + x_3 + x_4 = 7$$
$$x_1 + 2x_2 + 2x_3 - x_4 = 12$$
$$2x_1 + 4x_2 + 6x_4 = 4$$

The corresponding coefficient matrix would be:

$$\textbf{A} = \begin{bmatrix}
1 & 2 & 1 & 1 & 7\\
1 & 2 & 1 & -1 & 12\\
2 & 4 & 0 & 6 & 4
\end{bmatrix}$$

Using this form, you can do all the operations you could on normal equations: addition, scalar multiplication, and so on. As a first step, you want only a single coefficient of 1 in a column. This is known as *reduced row echelon form*. This allows you to solve for variables.

In an ideal world, the first step you would want to make the bottom two rows all equal to zero. By replacing the second row with the first row minus the second row, and subtracting twice the top row from the bottom row:

$$\begin{bmatrix}
1 & 2 & 1 & 1 & 7\\
0 & 0 & -1 & 2 & -5\\
0 & 0 & -2 & 4 & -10
\end{bmatrix}$$

Now, you want to make the -1 coefficient in the second row into a positive. You do this by multiplying that row by -1:

$$\begin{bmatrix}
1 & 2 & 1 & 1 & 7\\
0 & 0 & 1 & -2 & 5\\
0 & 0 & -2 & 4 & -10
\end{bmatrix}$$

To eliminate the bottom row, you add the bottom row to twice the value of the middle row.

$$\begin{bmatrix}
1 & 2 & 1 & 1 & 7\\
0 & 0 & 1 & -2 & 5\\
0 & 0 & 0 & 0 & 0
\end{bmatrix}$$

Now you replace the first row with the first row minus the second row:

$$\begin{bmatrix}
1 & 2 & 0 & 3 & 7\\
0 & 0 & 1 & -2 & 5\\
0 & 0 & 0 & 0 & 0
\end{bmatrix}=\text{rref}(\textbf{A})$$

Now the matrix is in reduced row echelon format. You know this is in reduced row echelon form, because all of your leading ones in each row are the *only* non-zero entry in their columns. They are called *pivot entries*. If you have a zeroed-out row, that has to be your last row. Now we can go back to linear equation form.

$$x_1 + 2x_2 + 3x_4 = 2$$
$$x_3 - 2x_4 = 5$$

The variables associated with each pivot entry are called *pivot variables*. Those that aren't are considered *free variables*. You can only solve for your pivot variables. You can only solve for your pivot variables. The reason why this equation has multiple solutions is because it is unconstrained.

$$x_1 = 2 -2x_2 - 3x_4$$
$$x_3 = 5 + 2x_4$$

This is as far as we can go to the system of linear equations. Hence,

$$\begin{bmatrix}
x_1 \\ x_2 \\ x_3\\ x_4\\
\end{bmatrix}=\text{rref}(\textbf{A})$$

***

### Another example

Here is an example of a solveable system of linear equations:

$$x + y + z = 3$$
$$x + 2y + 3z = 0$$
$$x + 3y + 4z = -2$$

The coefficient matrix is equal to: 

$$\begin{bmatrix}
1 & 1 & 1 & 3 \\
1 & 2 & 3 & 0 \\
1 & 3 & 4 & -2 \\
\end{bmatrix}$$

Now, we want to eliminate the first middle one, so we subtract first row from the second row, and replace that in the second row. We do the same for the third row.:

$$\begin{bmatrix}
1 & 1 & 1 & 3 \\
0 & 1 & 2 & -3 \\
0 & 2 & 3 & -5 \\
\end{bmatrix}$$

To zero the top row's middle one, we subtract the middle row from the top row and replace that in the top row, and do the same with 2 times the middle row for the bottom:

$$\begin{bmatrix}
1 & 0 & -1 & 6 \\
0 & 1 & 2 & -3 \\
0 & 0 & -1 & 1 \\
\end{bmatrix}$$

We want a pivot entry in the third row, so we multiply that row by -1.

$$\begin{bmatrix}
1 & 0 & -1 & 6 \\
0 & 1 & 2 & -3 \\
0 & 0 & 1 & -1 \\
\end{bmatrix}$$

So now, it's easy to zero out the middle and top row for the third column by adding the final row to the top row, and subtracting twice the bottom row from the middle row.

$$\begin{bmatrix}
1 & 0 & 0 & 5 \\
0 & 1 & 0 & -1 \\
0 & 0 & 1 & -1 \\
\end{bmatrix}$$

There are no free variables – every column has a pivot entry! Now, this gives us the values of $x = 5, y = -1, z = -1$.

***

In a system that has no solutions, finding reduced row echelon form makes one of the equations **false**. To illustrate this, consider the system of liar equations:

$$x_1 + 2x_2 + x_3 + x_4 = 8$$
$$x_1 + 2x_2 + 2x_3 - x_4 = 12$$
$$2x_1 + 4x_2 + 6x_4 = 4$$

This in a coefficient matrix is:

$$\begin{bmatrix}
1 & 2 & 1 & 1 & 8 \\
1 & 2 & 2 & -1 & 12 \\
2 & 4 & 0 & 6 & 4 \\
\end{bmatrix} = 
\begin{bmatrix}
1 & 2 & 1 & 1 & 8 \\
0 & 0 & 1 & -2 & 4 \\
0 & 0 & -2 & 4 & -12 \\
\end{bmatrix} = 
\begin{bmatrix}
1 & 2 & 0 & 3 & 4 \\
0 & 0 & 1 & -1 & 4 \\
0 & 0 & 0 & 0 & -4 \\
\end{bmatrix}$$

Since the last row has $0 = -4$, there are no ways to satisfy all three equations, meaning there are no solutions. Geometrically, this would mean that there are two non-intersecting lines, planes, or hyperplanes.

Whenever you find reduced row echelon format, there are three possible scenarios:

* $0 = a$, meaning there is no solution.
* All ones are in their own columns, meaning there is a unique solution.
* There are many solutions (unable to isolate the ones).