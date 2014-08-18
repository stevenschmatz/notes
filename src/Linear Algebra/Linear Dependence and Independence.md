\title{Linear Algebra: Linear Dependence and Independence}
\date{August 17, 2014}
\author{Steven Schmatz, University of Michigan: College of Engineering}

\maketitle
\begin{center} stevenschmatz@gmail.com \end{center} \pagebreak

<!---Content goes here-->

Introduction to linear independence
===================================

If two vectors are *linearly dependent*, any linear combination of the vectors will produce a scalar combination of one vector. For example,

$$\hat{\textbf{a}} = (2, 3)$$
$$\hat{\textbf{b}} = (4, 6)$$
$$c_1 \hat{\textbf{a}} + c_2 \hat{\textbf{b}} = c_1 \hat{\textbf{a}} + 2 c_2 \hat{\textbf{b}}$$
$$=(c_1 + 2c_2)\hat{\textbf{a}} = c_3\hat{\textbf{a}}$$

Hence $\hat{\textbf{a}}$ and $\hat{\textbf{b}}$ are linearly dependent.

Two vectors are linearly dependent if one vector in the set can be represented by a linear combination of another vector in that set.

Even a set of three vectors in $\mathbb{R}^2$ would be linearly dependent, because the third vector could be represented by a linear combination of the other two vectors *if they span the vector space*. The third vector would have to be outside of the $\mathbb{R}^2$ plane to be linearly independent.

A good *basis* for $\mathbb{R}^2$ would be two, linearly independent vectors. Hence, unit vectors $\hat{\textbf{i}}$ and $\hat{\textbf{j}}$ would be a good basis for $\mathbb{R}^2$.

More on linear independence
===========================

A set $s = \{v_1, v_2 \dots v_n\}$ is *linearly dependent* $\Longleftrightarrow$ $c_1v_1 + c_2v_2 + \dots c_nv_n = \textbf{0}$, for some $c_i$ where at least one is $\ne 0$. This means that one vector can be represented by the sum of the other vectors:

$$v_1 = a_2v_2 + a_3v_3 + \dots a_nv_n$$
$$0 = -1v_1 + a_2v_2 + a_3v_3 + \dots a_nv_n$$

Assuming $c_1 \ne 0$, you can divide both sides by $c_1$, yielding:

$$v_1 + \frac{c_2}{c_1}v_2 + \dots \frac{c_n}{c_1}= \textbf{0}$$
$$\frac{c_2}{c_1}v_2 + \dots \frac{c_n}{c_1}v_n = -v_1$$

Hence, if at least one of these constants is nonzero, you can represent $v_1$ as a linear combination of the other vectors.

Testing for linear independence
-------------------------------

If you have a linear combination of two vectors $c_1v_1 + c_2v_2 = \textbf{0}$:

* If $c_1$ or $c_2$ are nonzero $\rightarrow$ they are dependent.
* Else if $c_1$ *and* $c_2$ are both zero $\rightarrow$ they are independent.

For example:

$$v_1 = (2, 1), v_2 = (3, 2)$$
$$2c_1 + 3c_2 = 0$$
$$c_1 + 2c_2 = 0$$
$$c_1 + \frac{3}{2}c_2 = 0$$
$$\frac{1}{2}c_2 = 0$$
$$c_2 = 0 = c_1$$

Hence, these vectors are linearly independent.

For three vectors in $\mathbb{R}^2$:

$$v_1 = (2,1), v_2 = (3,2), v_3 = (1, 2)$$
$$2c_1 + 3c_2 + c_3 = 0$$
$$c_1 + 2c_2 + 2c_3 = 0$$

Picking a $c_3 = -1$:

$$2c_1 + 3c_2 - 1 = 0$$
$$2c_1 + 4c_2 - 4 = 0$$
$$-c_2 + 3 = 0$$
$$c_2 = 3$$
$$c_1 = -4, c_2 = 3, c_3 = -1$$

If you had three vectors in two-dimensional space, one of them has to be linearly dependent.

If you can show that you can satisfy $c_1v_1 + c_2v_2 = \textbf{0}$ with at least one nonzero constant, the vectors are linearly dependent.

Span and linear independence example
------------------------------------
1. Does the following set of vectors span $\mathbb{R}^3$?
2. Are the vectors linearly independent?

### Testing for span

$$s = \{v_1 = (1, -1, 2), v_2 = (2, 1, 3), v_3 = (-1, 0, 2)\}$$

Can $c_1v_1 + c_2v_2 + c_3v_3 = (a, b, c)$ for any $a, b, c$?

$$c_1 + 2c_2 - c_3 = a$$
$$-c_1 + c_2 = b$$
$$2c_1 + 3c_2 + 2c_3 = c$$

$$b+a = 3c_2 - c_3$$

Multiplying $a$ by -2:

$$c-2a = -c_2 + 4c_3$$

To eliminate the $-c_2$ term in the last equation, we multiply it by 3 and add it to the $b+a$ term.

$$11c_3 = 3c -5a + b$$
$$c_3 = \frac{1}{11}(3c-5a+b)$$
$$c_2 = \frac{1}{3}(b+a+c_3)$$
$$c_1 = a - 2c_2 + c_3$$

Hence, with any $a, b, c$, you can find $c_1, c_2, c_3$. They are then said to span $\mathbb{R}^3$.

### Testing for linear independence

$$c_1v_1 + c_2v_2 + c_3v_3 = \textbf{0}$$.

It's the same procedure as the last problem, but using $\textbf{0}$ instead of $a, b, c$.

$$c_3 = \frac{1}{11}(3(0) - 5(0) + 0) = 0$$
$$c_2 = \frac{1}{3}(0 + 0 + 0) = 0$$
$$c_1 = (0 -2(0)+0) = 0$$

Since all of these constants are equal to zero, these vectors are linearly independent.