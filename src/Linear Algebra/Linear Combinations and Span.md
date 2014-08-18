\title{Linear Algebra: Linear Combinations and Span}
\date{August 17, 2014}
\author{Steven Schmatz, University of Michigan: College of Engineering}

\maketitle
\begin{center} stevenschmatz@gmail.com \end{center} \pagebreak

<!---Content goes here-->

Linear combination
==================

A *linear combination* is a sum of vectors $v_1 \dots v_n \in \mathbb{R}^m$, with each multiplied by constants $c_1 \dots c_n \in \mathbb{R}$. For example, say we have two vectors:

$$\hat{\textbf{a}} = \hat{\textbf{i}} + 2 \hat{\textbf{j}}$$
$$\hat{\textbf{b}} = 3 \hat{\textbf{j}}$$

For example, here are a few valid linear combinations of these vectors.

$$0\hat{\textbf{a}} + 0\hat{\textbf{b}} = \textbf{0}$$
$$3\hat{\textbf{a}} + (-2)\hat{\textbf{b}} = 3\hat{\textbf{i}}$$

Span
----

The *span* is the set of all the vectors you can represent by adding and subtracting vectors.

$$\text{span}(v_1, v_2, \dots v_n) = \{c_1 v_1 + c_2 v_2 \dots c_n v_n \mid c_i \in \mathbb{R}\}$$

For example with vectors $\hat{\textbf{a}}$ and $\hat{\textbf{b}}$:

$$\text{span}(\hat{\textbf{a}}, \hat{\textbf{b}}) = \mathbb{R}^m$$

If two vectors are collinear, their span is **not** all of $\mathbb{R}^m$! It is $c \hat{\textbf{a}}$.

$$\text{span}(\textbf{0}) = \textbf{0}$$

The most familiar vectors that span the $\mathbb{R}^2$ vector space are the unit vectors $\hat{\textbf{i}}$ and $\hat{\textbf{j}}$. These form the *basis* of $\mathbb{R}^2$.