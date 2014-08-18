\title{Linear Algebra: Subspaces and the Basis for a Subspace}
\date{August 18, 2014}
\author{Steven Schmatz, University of Michigan: College of Engineering}

\maketitle
\begin{center} stevenschmatz@gmail.com \end{center} \pagebreak

<!---Content goes here-->

# Linear Subspaces

If a set $V$ is a *subspace* of $\mathbb{R}^n$:

1. $V$ contains $\textbf{0}.$
2. **Closure under scalar multiplication:** If $\hat{\textbf{x}}$ in $V$, then $c\hat{\textbf{x}}$ is in V.
3. **Closure under addition:** If $\hat{\textbf{a}}$ and $\hat{\textbf{b}}$ in $V$, then $\hat{\textbf{a}} + \hat{\textbf{b}}$ is in $V$.

For example, take the simple example $V = \{\textbf{0}\}$.

1. Does $V$ contain $\textbf{0}$? Yes.
2. Is there closure under scalar multiplication? $c*\textbf{0} = \textbf{0}$, so yes.
3. Is there closure under addition? $\textbf{0} + \textbf{0} = \textbf{0}$.

Hence, this trivial example *is* a subspace of $\mathbb{R}^m$.

***

For a less trivial example:

$$S = \{(x_1, x_2) \in \mathbb{R}^2 \mid x_1 \ge 0\}$$

Is $S$ a subspace of $\mathbb{R}^2$?

1. Does $V$ contain $\textbf{0}$? Yes.
2. Is there closure under scalar multiplication? **No**, because scalar multiplication by a negative number would not be a member of $S$!
3. Is there closure under addition? Yes, because all vectors have a positive $x_1$ value so the resultant vector would have a $x_1 \ge 0$.

Hence, since all three conditions are not true, this is not a subspace of $\mathbb{R}^2$.

***

If we define $U = \text{span}(v_1, v_2, v_3)$, is it a valid subspace of $\mathbb{R}^n$?

1. Does $U$ contain $\textbf{0}$? Yes.
$$0v_1 + 0v_2 + 0v_3 = \textbf{0}$$
2. Is there closure under scalar multiplication? Yes.
$$a*\hat{\textbf{x}} = ac_1v_1 = ac_2v_2 + ac_3v_3 = c_4v_1 + c_5v_2 + c_6v_3$$ 
3. Is there closure under addition? Yes.
$$\hat{\textbf{x}} + \hat{\textbf{y}} = (c_1 + d_1)v_1 + (c_2 + d_2)v_2 + (c_3 + d_3)v_3$$

Hence, $U$ is a valid subspace of $\mathbb{R}^n$

***

Is the set $U = \text{span}((1,1))$ a valid subspace of $\mathbb{R}^2$?

1. Does $U$ contain $\textbf{0}$? Yes.
$$0(1, 1) = (0, 0)$$
2. Is there closure under scalar multiplication? Yes.
$$\hat{\textbf{a}} = c(1, 1)$$
3. Is there closure under addition? Yes.
$$c_1(1, 1) + c_2(1,1) = (c_1 + c_2)(1, 1)$$

Hence, $U$ is a valid subspace of $\mathbb{R}$.

# Basis of a subspace

A *basis* of a subspace is a collection of linearly independent vectors, such that when you take the span of all those vectors, you can construct any of the vectors in the subspace.

$$S = \{v_1, v_2, \dots v_n\}$$
$$V = \text{span}(S) | S \text{ is linearly independent.}$$

In this case, $S$ is a basis for $V$.

A basis can be thought of a "minimum" set of vectors that spans the subspace.

Hence, the unit vectors $\hat{\textbf{i}} = (1, 0)$ and $\hat{\textbf{j}} = (0, 1)$ provide a basis for cartesian coordinates, because they span $\textbf{R}^2$ and they are linearly independent.